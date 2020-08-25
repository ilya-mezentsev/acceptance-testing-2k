export class Codes {
  private readonly defaultMessage: string = 'Unknown error';
  protected descriptionsToErrorMessage: Map<string, string> = new Map<string, string>([
    ['unable-to-decode-request', 'Unknown request error'],
    ['invalid-request-data', 'Entered data is invalid'],
    ['db-error', 'Internal service error']
  ]);

  public getMessageByDescription(description: string): string {
    if (this.descriptionsToErrorMessage.has(description)) {
      return this.descriptionsToErrorMessage.get(description);
    }

    return this.defaultMessage;
  }
}
