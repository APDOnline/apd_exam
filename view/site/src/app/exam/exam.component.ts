import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';

import 'rxjs/add/operator/switchMap';

import { Exam } from '../model/exam';
import {ExamService} from "./exam.service";
import {Book} from "../model/Book";
@Component({
  selector: 'app-exam',
  templateUrl: './exam.component.html',
  styleUrls: ['./exam.component.scss']
})
export class ExamComponent implements OnInit {

  exam: Exam[] = [];
  book: Book;
  questionCount: number;

  constructor(private route: ActivatedRoute, private examService: ExamService) { }

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => {
      return this.examService.getExam(+params['examId'])
    })
      .subscribe(exam => {
        this.book = exam.book;
        this.questionCount = exam.question ? exam.question.length : 0;
        this.exam.push(exam);
      })
  }

  addQuestion(question) {
    this.exam[0].question = this.exam[0].question || [];
    this.exam[0].question.push(question);
    this.questionCount = this.exam[0].question.length;
  }

}
