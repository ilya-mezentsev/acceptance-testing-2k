import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {StorageService} from "../services/storage/storage.service";
import {ErrorHandlerService} from "../../services/errors/error-handler.service";
import {Object} from "../types/types";

@Component({
  selector: 'app-edit-object',
  templateUrl: './edit-object.component.html',
  styleUrls: ['./edit-object.component.scss']
})
export class EditObjectComponent implements OnInit {
  private objectHash = '';
  private currentObject: Object;

  constructor(
    private readonly router: Router,
    private readonly route: ActivatedRoute,
    private readonly storage: StorageService,
    private readonly errorHandler: ErrorHandlerService,
  ) { }

  public get currentObjectName(): string {
    return this.currentObject?.name;
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.objectHash = params.get('object_hash');
      this.setCurrentObjectName();
    });
  }

  private setCurrentObjectName(): void {
    if (this.storage.hasObjects()) {
      for (const object of this.storage.objects) {
        if (object.hash === this.objectHash) {
          this.currentObject = object;
          return;
        }
      }
    }

    this.router.navigate(['/admin'])
      .catch(err => this.errorHandler.handle(err));
  }
}
