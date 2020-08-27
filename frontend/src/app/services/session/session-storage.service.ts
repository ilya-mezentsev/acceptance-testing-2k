import { Injectable } from '@angular/core';

type Session = {account_hash: string, login: string};

@Injectable({
  providedIn: 'root'
})
export class SessionStorageService {
  private session: Session;
  private readonly sessionKey = 'AAT-Session';

  public deleteSession(): void {
    this.saveSession(null);
  }

  public hasSession(): boolean {
    const hasSession = !!this.session?.account_hash && !!this.session?.login;
    if (!hasSession) {
      this.tryRestoreSession();
    }

    return !!this.session;
  }

  private tryRestoreSession(): void {
    this.session = JSON.parse(localStorage.getItem(this.sessionKey));
  }

  public saveSession(session: Session): void {
    localStorage.setItem(this.sessionKey, JSON.stringify(session));
    this.session = session;
  }

  public getSessionId(): string {
    return this.session?.account_hash;
  }

  public getSessionLogin(): string {
    return this.session?.login;
  }
}
