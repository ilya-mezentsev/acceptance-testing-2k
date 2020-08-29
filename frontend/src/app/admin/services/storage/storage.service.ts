import { Injectable } from '@angular/core';
import {TestCommand, TestObject} from '../../types/types';

@Injectable({
  providedIn: 'root'
})
export class StorageService {
  private readonly testObjectsKey = 'test-objects';
  private readonly testCommandsKey = 'test-commands';
  private readonly storage: Map<string, any> = new Map<string, any>();

  public invalidateObjects(): void {
    this.storage.delete(this.testObjectsKey);
  }

  public hasObjects(): boolean {
    return this.storage.has(this.testObjectsKey);
  }

  public get objects(): TestObject[] {
    if (!this.storage.has(this.testObjectsKey)) {
      return [];
    }

    return this.storage.get(this.testObjectsKey);
  }

  public set objects(objects: TestObject[]) {
    this.storage.set(this.testObjectsKey, objects);
  }

  public hasCommands(): boolean {
    return this.storage.has(this.testCommandsKey);
  }

  public get commands(): TestCommand[] {
    if (!this.storage.has(this.testCommandsKey)) {
      return [];
    }

    return this.storage.get(this.testCommandsKey);
  }

  public set commands(commands: TestCommand[]) {
    this.storage.set(this.testCommandsKey, commands);
  }

  public invalidateCommands(): void {
    this.storage.delete(this.testCommandsKey);
  }
}
