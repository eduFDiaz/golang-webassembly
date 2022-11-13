import { TestBed } from '@angular/core/testing';

import { WasmService } from './wasm.service';

describe('DataServiceService', () => {
  let service: WasmService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WasmService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
