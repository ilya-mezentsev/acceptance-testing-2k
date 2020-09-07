import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {HashService} from "../../services/hash/hash.service";

@Component({
  selector: 'app-input-number',
  templateUrl: './input-number.component.html',
  styleUrls: ['./input-number.component.scss']
})
export class InputNumberComponent implements OnInit {
  public id = '';
  @Input() public value = 0;
  @Input() public min = Number.MIN_SAFE_INTEGER;
  @Input() public max = Number.MAX_SAFE_INTEGER;
  @Input() public readonly label = '';
  @Output() public valueEmitter = new EventEmitter<number>();

  constructor(
    private readonly hashService: HashService,
  ) { }

  public valueChanged(): void {
    this.valueEmitter.emit(this.value);
  }

  ngOnInit(): void {
    this.id = this.hashService.getRandomHash();
  }
}
