// Copyright 2023 The Kubermatic Kubernetes Platform contributors.
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
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ClusterBackup} from '@app/shared/entity/backup';

@Component({
  selector: 'km-delete-backup-dialog',
  templateUrl: './template.html',
})
export class DeleteBackupDialogComponent implements OnInit {
  verificationInput = '';
  backups: ClusterBackup[] = [];

  constructor(
    @Inject(MAT_DIALOG_DATA) private readonly _config: DeleteBackupDialogComponent,
    private readonly _dialogRef: MatDialogRef<DeleteBackupDialogComponent>
  ) {}

  ngOnInit(): void {
    this.backups = this._config.backups;
  }

  onEnterKeyDown(): void {
    if (!this.isNameVerified()) {
      return;
    }
    this._dialogRef.close(true);
  }

  isNameVerified(): boolean {
    if (this.backups.length > 1) {
      return this.verificationInput === 'yes';
    }
    return this.verificationInput === this.backups[0]?.name;
  }
}