import { Component } from '@angular/core';

@Component({
  selector: 'app-blog-editor',
  templateUrl: './create.html',
  styleUrls: ['./create.scss'],
})
export class Create {
  title = '';
  content = '';
  tags: string[] = [];
  tagInput = '';
  isPublishing = false;
  showPreview = false;

  handleTagAdd() {
    const normalized = this.tagInput.trim().toLowerCase();
    if (!normalized || this.tags.includes(normalized)) return;
    this.tags.push(normalized);
    this.tagInput = '';
  }

  removeTag(index: number) {
    this.tags.splice(index, 1);
  }

  handlePublish() {
    if (!this.title.trim() || !this.content.trim()) return;
    this.isPublishing = true;

    setTimeout(() => {
      this.isPublishing = false;
      this.showPreview = true;
    }, 700);
  }

  handleClear() {
    this.title = '';
    this.content = '';
    this.tags = [];
    this.tagInput = '';
    this.showPreview = false;
  }
}
