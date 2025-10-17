import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { BlogService, Blog } from '../../../core/application/blog/blog.service';
import { Location } from '@angular/common';
import { CommentBoxComponent } from '../../../ui/comment/comment-box/comment-box';

@Component({
  selector: 'app-detail',
  standalone: true,
  imports: [CommonModule, CommentBoxComponent],
  templateUrl: './detail.html',
})
export class DetailComponent {
  blog?: Blog;

  constructor(
    private route: ActivatedRoute,
    private blogService: BlogService,
    private location: Location
  ) {}

  goBack(): void {
    this.location.back();
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.blog = this.blogService.getBlogById(id);
    }
  }
}