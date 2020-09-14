import {
  Component,
  Input,
  OnInit,
  Output,
  EventEmitter,
  ViewChild,
  ElementRef
} from '@angular/core';

@Component({
  selector: 'app-input-file',
  templateUrl: './input-file.component.html',
  styleUrls: ['./input-file.component.scss']
})
export class InputFileComponent implements OnInit {
  @ViewChild('fileInput') fileInput: ElementRef;
  @ViewChild('filenameInput') filenameInput: ElementRef;
  @Input() public label = '';
  @Input() public accept = '';
  @Input() public resetEvent = new EventEmitter();
  @Output() public fileChanged = new EventEmitter<any>();

  constructor() { }

  public reset(): void {
    this.fileInput.nativeElement.value = null;
    this.filenameInput.nativeElement.value = '';
    this.fileChanged.emit();
  }

  ngOnInit(): void {
    this.resetEvent.subscribe(() => this.reset());
  }
}
