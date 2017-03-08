import { Component, OnInit } from '@angular/core';

import {MenuItem} from 'primeng/primeng';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  private menuItem: MenuItem;

  ngOnInit() {
    this.menuItem = [
      {label: 'Books', icon: 'fa-book', routerLink: ["/book"]},
      {label: 'Exams', icon: 'fa-list', routerLink: ["/examList"]}
    ]
  }

}
