import {Component, Input, OnInit} from '@angular/core';
import {StorageService} from '../../../services/storage/storage.service';
import {Router} from '@angular/router';
import {ErrorHandlerService} from '../../../../services/errors/error-handler.service';
import {GeneralUpdateRequest, TestCommandRecord, TestObject} from '../../../types/types';

@Component({
  selector: 'app-commands-table',
  templateUrl: './commands-table.component.html',
  styleUrls: ['./commands-table.component.scss']
})
export class CommandsTableComponent implements OnInit {
  @Input() private generalUpdateRequest: GeneralUpdateRequest = {
    command_hashes: []
  };

  constructor(
    private readonly router: Router,
    private readonly storage: StorageService,
    private readonly errorHandler: ErrorHandlerService,
  ) { }

  public hasObjects(): boolean {
    return this.storage.objects.length > 0;
  }

  public getObjects(): TestObject[] {
    return this.storage.objects.filter(
      o => !!this.storage.commands.filter(c => c.object_hash === o.hash).length
    );
  }

  public getObjectCommands(objectHash: string): TestCommandRecord[] {
    return this.storage.commands.filter(c => c.object_hash === objectHash);
  }

  public getChosenHashes(): string[] {
    return this.generalUpdateRequest.command_hashes;
  }

  public addHash(commandHash: string): void {
    this.generalUpdateRequest.command_hashes.push(commandHash);
  }

  public removeHash(commandHash: string): void {
    this.generalUpdateRequest.command_hashes = this.generalUpdateRequest.command_hashes
      .filter(h => h !== commandHash);
  }

  public addAllHashes(): void {
    this.generalUpdateRequest.command_hashes = this
      .getObjects()
      .map(o => o.hash)
      .map(objectHash => this.getObjectCommands(objectHash).map(c => c.hash))
      .reduce((acc, commandsHashes) => acc.concat(commandsHashes), []);
  }

  public removeAllHashes(): void {
    this.generalUpdateRequest.command_hashes = [];
  }

  public allObjectCommandsIsChosen(objectHash: string): boolean {
    return this.getObjectCommands(objectHash)
      .map(c => c.hash)
      .every(h => this.generalUpdateRequest.command_hashes.includes(h));
  }

  public addAllObjectCommands(objectHash: string): void {
    this.generalUpdateRequest.command_hashes
      .push(...this.getObjectCommands(objectHash).map(c => c.hash));
  }

  public removeAllObjectCommands(objectHash: string): void {
    const objectCommandsHashes = this.getObjectCommands(objectHash).map(c => c.hash);

    this.generalUpdateRequest.command_hashes = this.generalUpdateRequest.command_hashes
      .filter(h => !objectCommandsHashes.includes(h));
  }

  ngOnInit(): void {
    if (!this.storage.hasObjects()) {
      this.router.navigate(['/admin'])
        .catch(err => this.errorHandler.handle(err));
    }
  }
}
