import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

export type AddFunctionInterface = (val1: string, val2: string, result: string) => number;

export interface WasmSuite {
  addFunction: AddFunctionInterface;
}

declare const window: any;

@Injectable({
  providedIn: 'root'
})
export class WasmService {

  public ready = new BehaviorSubject<boolean>(false);

  private Suite!: WasmSuite;

  constructor() {
    // Init wasm, then update the ready state
    this.init().then(_ => {
      this.ready.next(true);
    });
  }

  private async init() {
    let go = new Go();
    let mod, inst;

    if (!WebAssembly.instantiateStreaming) {
      // polyfill
      WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }

    return WebAssembly.instantiateStreaming(
      fetch('assets/wasm/lib.wasm'),
      go.importObject
    ).then(
      async result => {
        mod = result.module;
        inst = result.instance;
        await go.run(inst);
      }
    );

  }
  public callAdd(a: string, b: string) {
    return window.add(a,b);
  }

  public factorial_js(a: number) :number{
    if (a <= 1){
      return a;
    }
    return a*this.factorial_js(a-1);
  }
}
