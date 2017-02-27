import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {Exam} from "../model/Exam";
import 'rxjs/add/operator/toPromise'

@Injectable()
export class ExamService {

  getExamUrl: string = "http://localhost:8112/api/exam/get/";

  constructor(private http: Http) {
  }

  getExam(examId: number): Promise<Exam> {
    return this.http.get(this.getExamUrl + examId)
      .toPromise()
      .then(response => response.json() as Exam)
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }

}
