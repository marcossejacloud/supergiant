import { Component, OnInit, ViewEncapsulation, OnDestroy } from '@angular/core';
import { CloudAccountModel } from '../cloud-accounts.model';
import { Subscription } from 'rxjs/Subscription';
import { Supergiant } from '../../../shared/supergiant/supergiant.service';
import { Notifications } from '../../../shared/notifications/notifications.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-new-cloud-account',
  templateUrl: './new-cloud-account.component.html',
  styleUrls: ['./new-cloud-account.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class NewCloudAccountComponent implements OnInit, OnDestroy {
  private providersObj = new CloudAccountModel;
  private subscriptions = new Subscription();
  private providers = this.providersObj.providers;
  private model: any;
  public schema: any;

  constructor(
    private supergiant: Supergiant,
    private notifications: Notifications,
    private router: Router,
  ) { }

  ngOnInit() {
  }

  ngOnDestroy() {
    this.subscriptions.unsubscribe();
  }

  back() {
    this.model = null;
    this.schema = null;
  }

  create(model) {
    // TODO: find a better way to do this...
    if (model.provider === 'gce') {
      model.credentials = JSON.parse(model.credentials.service_account_key);
    }
    this.subscriptions.add(this.supergiant.CloudAccounts.create(model).subscribe(
      (data) => {
        this.success(model);
        this.router.navigate(['/system/cloud-accounts']);
      },
      (err) => { this.error(model, err); }));
  }

  success(model) {
    this.notifications.display(
      'success',
      'Kube: ' + model.name,
      'Created...',
    );
  }

  error(model, data) {
    this.notifications.display(
      'error',
      'Kube: ' + model.name,
      'Error:' + data.statusText);
  }
}
