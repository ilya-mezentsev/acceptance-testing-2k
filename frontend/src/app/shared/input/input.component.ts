import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';
import {Md5} from 'ts-md5';
import {HashService} from '../../services/hash/hash.service';

@Component({
  selector: 'app-input',
  templateUrl: './input.component.html',
  styleUrls: ['./input.component.scss']
})
export class InputComponent implements OnInit {
  public id = '';
  @Input() public value = '';
  @Input() public readonly label = '';
  @Input() public readonly type = '';
  @Output() public valueEmitter = new EventEmitter<string>();

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
