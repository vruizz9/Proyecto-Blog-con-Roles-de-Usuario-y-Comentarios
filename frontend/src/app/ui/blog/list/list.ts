// src/app/ui/blog/list/list.ts
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { PaginationComponent } from '../../../shared/componentes/pagination/pagination';
import { BlogService, Blog } from '../../../core/application/blog/blog.service';

@Component({
  selector: 'app-list',
  standalone: true,
  imports: [CommonModule, RouterModule, PaginationComponent],
  templateUrl: './list.html',
})
export class List {
  blogs: Blog[] = [];
  currentPage = 1;
  itemsPerPage = 5;

  constructor(private blogService: BlogService) {
    this.blogs = this.blogService.getBlogs();
  }

  get paginatedBlogs(): Blog[] {
    const start = (this.currentPage - 1) * this.itemsPerPage;
    return this.blogs.slice(start, start + this.itemsPerPage);
  }

  onPageChange(page: number) {
    this.currentPage = page;
  }

  formatDate(iso: string) {
    const d = new Date(iso);
    return d.toLocaleDateString(undefined, {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  }
}