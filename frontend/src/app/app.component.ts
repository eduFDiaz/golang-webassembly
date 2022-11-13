import { Component } from '@angular/core';
import { WasmService } from './wasm.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  value1 = '1';
  value2 = '2';
  result = '';

  constructor (private wasmService: WasmService){

  }

  public add(){
    console.log(this.wasmService.callAdd(this.value1,this.value2));

    console.log(this.wasmService.factorial_js(5));
  }
}
