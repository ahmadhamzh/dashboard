// Copyright 2020 The Kubermatic Kubernetes Platform contributors.
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

import {ComponentFixture, TestBed, waitForAsync} from '@angular/core/testing';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {BrowserModule} from '@angular/platform-browser';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import {CoreModule} from '@core/module';
import {MLAService} from '@core/services/mla';
import {NotificationService} from '@core/services/notification';
import {SharedModule} from '@shared/module';
import {fakeDigitaloceanCluster} from '@test/data/cluster';
import {fakeAlertmanagerConfig} from '@test/data/mla';
import {fakeProject} from '@test/data/project';
import {asyncData} from '@test/services/cluster-mock';
import {MatDialogRefMock} from '@test/services/mat-dialog-ref-mock';
import {MonacoEditorModule, NGX_MONACO_EDITOR_CONFIG} from 'ngx-monaco-editor-v2';
import {AlertmanagerConfigDialog} from './component';

declare let monaco: any;

describe('AlertmanagerConfigDialog', () => {
  let fixture: ComponentFixture<AlertmanagerConfigDialog>;
  let component: AlertmanagerConfigDialog;
  let putAlertmanagerConfigSpy: jest.Mock;

  beforeEach(waitForAsync(() => {
    const mlaMock = {
      putAlertmanagerConfig: jest.fn(),
      refreshAlertmanagerConfig: () => {},
    };
    putAlertmanagerConfigSpy = mlaMock.putAlertmanagerConfig.mockReturnValue(asyncData(fakeAlertmanagerConfig()));

    TestBed.configureTestingModule({
      imports: [BrowserModule, NoopAnimationsModule, SharedModule, CoreModule, MonacoEditorModule],
      declarations: [AlertmanagerConfigDialog],
      providers: [
        {provide: MatDialogRef, useClass: MatDialogRefMock},
        {provide: MLAService, useValue: mlaMock},
        {
          provide: MAT_DIALOG_DATA,
          useValue: {
            title: '',
            projectId: '',
            cluster: {},
            confirmLabel: '',
          },
        },
        NotificationService,
        {provide: NGX_MONACO_EDITOR_CONFIG, useValue: {onMonacoLoad: () => (monaco = (window as any).monaco)}},
      ],
      teardown: {destroyAfterEach: false},
    }).compileComponents();
  }));

  describe('Edit Alertmanager Config Dialog', () => {
    beforeEach(() => {
      fixture = TestBed.createComponent(AlertmanagerConfigDialog);
      component = fixture.componentInstance;

      component.data = {
        title: 'Edit Alertmanager Config',
        projectId: fakeProject().id,
        cluster: fakeDigitaloceanCluster(),
        confirmLabel: 'Edit',
        alertmanagerConfig: fakeAlertmanagerConfig(),
      };
      fixture.detectChanges();
    });

    it('should create the edit alertmanager config dialog', waitForAsync(() => {
      expect(component).toBeTruthy();
    }));

    it('should have correct title: edit', () => {
      expect(document.body.querySelector('km-dialog-title').textContent).toContain('Edit Alertmanager Config');
    });

    it('should have correct button text: edit', () => {
      expect(document.body.querySelector('#km-alertmanager-config-dialog-btn').textContent).toContain('Edit');
    });

    xit('should call patchAlertmanagerConfig()', () => {
      // component.save();
      fixture.detectChanges();
      expect(putAlertmanagerConfigSpy).toHaveBeenCalled();
    });
  });
});
