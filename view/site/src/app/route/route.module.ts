import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BookComponent} from "../book/book.component";
import {QuestionComponent} from "../question/question.component";
import {ExamComponent} from "../exam/exam.component";

const routes: Routes = [
  { path: 'book', component: BookComponent },
  { path: 'question/:bookName/:bookId', component: QuestionComponent },
  { path: 'exam/:examId', component: ExamComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class RouteModule { }
