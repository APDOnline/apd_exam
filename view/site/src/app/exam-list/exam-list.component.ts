import { Component, OnInit } from '@angular/core';
import {Exam} from "../model/Exam";
import {ExamListService} from "./exam-list.service";

@Component({
  selector: 'app-exam-list',
  templateUrl: './exam-list.component.html',
  styleUrls: ['./exam-list.component.scss']
})
export class ExamListComponent implements OnInit {

  exams: Exam[];
  error: any;

  constructor(private examListService: ExamListService) { }

  ngOnInit() {
    this.examListService.getList()
      .then(list => {
        this.exams = list;
      })
      .catch(error => {
        console.error(error);
        this.error = error;
      })

  }

}
