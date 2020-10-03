import { Injectable } from '@angular/core';
import {TestCommandRecord, TestObject} from '../../types/types';
import {RadioService} from '../../../services/radio/radio.service';
import {InvalidateStorage} from '../../../services/radio/const';

@Injectable({
  providedIn: 'root'
})
export class StorageService {
  private readonly testObjectsKey = 'test-objects';
  private readonly testCommandsKey = 'test-commands';
  private readonly allStorageKeys: string[] = [
    this.testObjectsKey,
    this.testCommandsKey,
  ];
  private readonly storage: Map<string, any> = new Map<string, any>();

  constructor(
    private readonly radio: RadioService
  ) {}

  // invalidate all storage data
  public invalidate(): void {
    for (const storageKey of this.allStorageKeys) {
      this.storage.delete(storageKey);
    }
  }

  public invalidateObjects(): void {
    this.storage.delete(this.testObjectsKey);

    this.radio.emit(InvalidateStorage);
  }

  public hasObjects(): boolean {
    return this.storage.has(this.testObjectsKey);
  }

  public get objects(): TestObject[] {
    if (!this.storage.has(this.testObjectsKey)) {
      return [];
    }

    return this.storage.get(this.testObjectsKey) || [];
  }

  public set objects(objects: TestObject[]) {
    this.storage.set(this.testObjectsKey, objects);
  }

  public hasCommands(): boolean {
    return this.storage.has(this.testCommandsKey);
  }

  public get commands(): TestCommandRecord[] {
    if (!this.storage.has(this.testCommandsKey)) {
      return [];
    }

    return this.storage.get(this.testCommandsKey) || [];
  }

  public set commands(commands: TestCommandRecord[]) {
    this.storage.set(this.testCommandsKey, commands);
  }

  public invalidateCommands(): void {
    this.storage.delete(this.testCommandsKey);

    this.radio.emit(InvalidateStorage);
  }
}
