// Copyright 2022 The Kubermatic Kubernetes Platform contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ExternalMachineDeploymentService} from '@core/services/external-machine-deployment';
import {NotificationService} from '@core/services/notification';
import {
  ExternalMachineDeployment,
  ExternalMachineDeploymentPatch,
  ExternalMachineDeploymentSpecPatch,
} from '@shared/entity/external-machine-deployment';
import {BaseFormValidator} from '@shared/validators/base-form.validator';
import {Observable} from 'rxjs';
import {filter, take, takeUntil} from 'rxjs/operators';
import {MasterVersion} from '@shared/entity/cluster';
import {minor, major} from 'semver';

class UpdateExternalClusterMachineDeploymentDialogData {
  projectID: string;
  clusterID: string;
  machineDeployment: ExternalMachineDeployment;
  kubeletVersion: string;
  replicas: number;
}

enum Controls {
  Replicas = 'replicas',
  KubeletVersion = 'version',
}

export enum KubeletVersionState {
  Empty = 'No Kubelet Versions Available',
  Loading = 'Loading...',
  Ready = 'Kubelet Version',
}

@Component({
  selector: 'km-update-external-cluster-machine-deployment-dialog',
  templateUrl: './template.html',
  styleUrls: ['./style.scss'],
  standalone: false,
})
export class UpdateExternalClusterMachineDeploymentDialogComponent extends BaseFormValidator implements OnInit {
  readonly Controls = Controls;
  disableReplicaControl: boolean;
  versions: string[] = [];
  versionsLabel = KubeletVersionState.Loading;
  upgradesAvailable: boolean;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: UpdateExternalClusterMachineDeploymentDialogData,
    private readonly _dialogRef: MatDialogRef<UpdateExternalClusterMachineDeploymentDialogComponent>,
    private readonly _notificationService: NotificationService,
    private readonly _externalMachineDeploymentService: ExternalMachineDeploymentService,
    private readonly _builder: FormBuilder
  ) {
    super();
  }

  ngOnInit(): void {
    this._initForm();
    this._initSubscriptions();
  }

  getObservable(): Observable<ExternalMachineDeployment> {
    const {projectID, clusterID, machineDeployment} = this.data;
    const patchModel = this._getExternalMachineDeploymentPatchModel(!!machineDeployment.cloud.eks);

    return this._externalMachineDeploymentService
      .patchExternalMachineDeployment(projectID, clusterID, machineDeployment.id, patchModel)
      .pipe(take(1));
  }

  onNext(md: ExternalMachineDeployment): void {
    this._notificationService.success(`Updated the ${md.name} machine deployment`);
    this._dialogRef.close();
  }

  private _initForm(): void {
    this.form = this._builder.group({
      [Controls.Replicas]: this._builder.control(this.data.replicas),
      [Controls.KubeletVersion]: this._builder.control(''),
    });
  }

  private _initSubscriptions(): void {
    const initialValue = this.form.value;
    this.form
      .get(Controls.Replicas)
      .valueChanges.pipe(takeUntil(this._unsubscribe))
      .subscribe((value: number) => {
        if (initialValue[Controls.Replicas] === value) {
          this.form.get(Controls.KubeletVersion).enable({emitEvent: false});
          this.form.get(Controls.KubeletVersion).clearValidators();
        } else {
          this.form.get(Controls.KubeletVersion).disable({emitEvent: false});
        }
      });

    this.form
      .get(Controls.KubeletVersion)
      .valueChanges.pipe(takeUntil(this._unsubscribe))
      .pipe(filter(_ => !!this.upgradesAvailable))
      .subscribe((value: string) => {
        this.disableReplicaControl = this.data.kubeletVersion !== value;
      });

    this._externalMachineDeploymentService
      .machineDeploymentUpgrades(this.data.projectID, this.data.clusterID, this.data.machineDeployment?.id)
      .pipe(take(1))
      .subscribe((upgrades: MasterVersion[]) => {
        this.upgradesAvailable = !!upgrades?.length;
        this.versionsLabel = this.upgradesAvailable ? KubeletVersionState.Ready : KubeletVersionState.Empty;
        if (this.upgradesAvailable) {
          this._setDefaultVersion(upgrades);
        }
      });
  }

  private _setDefaultVersion(upgrades: MasterVersion[]): void {
    this.versions = upgrades.map(upgrade => upgrade.version);
    const kubeletVersion = this.data.kubeletVersion;

    if (this.versions.includes(kubeletVersion)) {
      this.form.get(Controls.KubeletVersion).setValue(kubeletVersion);
      return;
    }
  }

  private _getExternalMachineDeploymentPatchModel(isEKSProvider: boolean): ExternalMachineDeploymentPatch {
    const kubeletVersion = this.form.get(Controls.KubeletVersion).value;

    const patchModel = {
      spec: {
        replicas: this.form.get(Controls.Replicas).value,
        template: {
          versions: {
            kubelet: kubeletVersion,
          },
        },
      } as ExternalMachineDeploymentSpecPatch,
    };

    // In EKS we shouldn't send the patch with the full version, only the major and minor versions.
    if (isEKSProvider && kubeletVersion) {
      patchModel.spec.template.versions.kubelet = `${major(kubeletVersion)}.${minor(kubeletVersion)}`;
    }
    return patchModel;
  }
}
