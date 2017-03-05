import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/switchMap';

import { Exam } from '../model/Exam';
import {ExamService} from "./exam.service";
import {Book} from "../model/Book";


@Component({
  selector: 'app-exam',
  templateUrl: './exam.component.html',
  styleUrls: ['./exam.component.scss']
})
export class ExamComponent implements OnInit {

  exam: Exam;
  book: Book;
  questionCount: number;
  examQuestionTabName: string;

  constructor(private route: ActivatedRoute, private examService: ExamService) { }

  ngOnInit() {
    this.updateExamQuestionTitle(0);
    this.exam = new Exam;
    this.route.params
      .switchMap((params: Params) => {
      return this.examService.getExam(+params['examId'])
    })
      .subscribe(exam => {
        this.book = exam.book;
        this.exam = exam;
        console.error(this.exam);
        this.updateExamQuestionTitle(this.exam.questions ? this.exam.questions.length : 0);
      })
  }

  addQuestion(question) {
    this.exam.questions = this.exam.questions || [];
    this.exam.questions.push(question);
    this.updateExamQuestionTitle(this.exam.questions.length);


    let newExam = this.examService.addQuestion(this.exam.id, question);
  }

  updateExamQuestionTitle(count: number) {
    this.examQuestionTabName = `Exam Questions (${count})`;
  }


}
