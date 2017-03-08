import { TestBed, inject } from '@angular/core/testing';
import { ExamListService } from './exam-list.service';

describe('ExamListService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ExamListService]
    });
  });

  it('should ...', inject([ExamListService], (service: ExamListService) => {
    expect(service).toBeTruthy();
  }));
});
