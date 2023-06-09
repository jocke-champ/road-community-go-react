// posts-list.component.ts
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-posts-list',
  templateUrl: './posts-list.component.html',
  styleUrls: ['./posts-list.component.css']
})
export class PostsListComponent implements OnInit {
  posts: any[] = [];

  ngOnInit() {
    // Fetch posts from your backend here
  }
}
