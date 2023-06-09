// new-post.component.ts
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PostService } from '../services/post.service';  // Assuming you have a PostService for making HTTP requests
import { AuthService } from '../services/auth.service';  // Assuming you have an AuthService for user management

@Component({
  selector: 'app-new-post',
  templateUrl: './new-post.component.html',
  styleUrls: ['./new-post.component.css']
})
export class NewPostComponent {
  postForm: FormGroup;

  constructor(private fb: FormBuilder, private postService: PostService, private authService: AuthService) {
    this.postForm = this.fb.group({
      title: ['', Validators.required],
      body: ['', Validators.required],
    });
  }

  submitPost() {
    if (!this.postForm.valid) {
      return;
    }

    const newPost = {
      title: this.postForm.get('title').value,
      body: this.postForm.get('body').value,
      author: this.authService.currentUser.id,  // Get the current user's ID from the AuthService
    };

    this.postService.submitPost(newPost.title, newPost.body, newPost.author)
      .subscribe(
        response => {
          console.log('Post submitted successfully', response);
          this.postForm.reset();
        },
        error => console.error('There was an error submitting the post', error)
      );
  }
}
