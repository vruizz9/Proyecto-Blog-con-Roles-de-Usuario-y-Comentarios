import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CommentService, Comment } from '../../../core/application/comment/comment.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-comment-box',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './comment-box.html',
  styleUrls: ['./comment-box.scss']
})
export class CommentBoxComponent {
  @Input() blogId!: string;

  comments: Comment[] = [];
  newComment: string = '';

  users: any[] = [];
  currentUser: any = null;

  constructor(
    private commentService: CommentService,
    private http: HttpClient
  ) {}

  ngOnInit(): void {
    this.loadMockUsers();
    this.loadComments();
  }
  loadMockUsers(): void {
    this.http.get<any[]>('http://localhost:3000/users').subscribe(users => {
      this.users = users;

      this.currentUser = this.users.find(u => u.username === 'Sandra') || this.users[0];
    });
  }

  loadComments(): void {
    if (!this.blogId) return;
    this.commentService.getCommentsByBlogId(this.blogId).subscribe(data => {
      this.comments = data;
    });
  }

  addComment(): void {
    if (!this.newComment.trim() || !this.currentUser) return;

    const newCommentObj: Comment = {
      id: Math.random().toString(36).substring(2, 9),
      blogId: this.blogId,
      userId: this.currentUser.id,
      content: this.newComment.trim()
    };

    this.commentService.addComment(newCommentObj).subscribe(() => {
      this.comments.push(newCommentObj);
      this.newComment = '';
    });
  }

  getUsername(userId: string): string {
    const user = this.users.find(u => u.id === userId);
    return user ? user.username : 'Usuario desconocido';
  }
}