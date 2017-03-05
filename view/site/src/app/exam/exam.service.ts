import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {Exam} from "../model/Exam";
import 'rxjs/add/operator/toPromise'
import {Question} from "../model/Question";

@Injectable()
export class ExamService {

  getExamUrl: string = "http://localhost:8112/api/exam/get/";
  addQuestionUrl: string = "http://localhost:8112/api/exam/updateQuestion";

  constructor(private http: Http) {
  }

  getExam(examId: number): Promise<Exam> {
    return this.http.get(this.getExamUrl + examId)
      .toPromise()
      .then(response => response.json() as Exam)
      .catch(this.handleError)
  }

  addQuestion(examId: number, question: Question): Promise<Exam> {
    let data = {
      exam_id: examId,
      question_id: question.id

    };
    return this.http.post(this.addQuestionUrl, JSON.stringify(data))
      .toPromise()
      .then(response => response.json() as Exam)
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }

}
