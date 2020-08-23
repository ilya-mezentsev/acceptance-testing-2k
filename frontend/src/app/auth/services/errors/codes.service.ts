import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CodesService {
  private readonly defaultMessage: string = 'Unknown error';
  private readonly descriptionsToErrorMessage: Map<string, string> = new Map<string, string>(
    [
      ['login-already-exists', 'Entered login already exists'],
      ['account-does-not-exists', 'Account does not exists']
    ]
  );

  public getMessageByDescription(description: string): string {
    if (this.descriptionsToErrorMessage.has(description)) {
      return this.descriptionsToErrorMessage.get(description);
    }

    return this.defaultMessage;
  }
}
