// components/CreatePost.jsx
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query'; // ✅ Fixed import
import { useCreatePost } from '../hooks/useCreatePost';
import { fetchPosts } from '../hooks/Post';
import "../styles/posts.css"
export default function CreatePost({ onSuccess }) {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [tags, setTags] = useState([]);
  const [tagInput, setTagInput] = useState('');
  const [error, setError] = useState('');

  // ✅ Fixed useQuery usage
  const { data: posts, isLoading } = useQuery({
    queryFn: () => fetchPosts(),
    queryKey: ["posts"],
    // Optional: Don't refetch on window focus
    refetchOnWindowFocus: false,
  });

  // ✅ Use the mutation hook
  const { mutate: createPost, isPending } = useCreatePost();

  const handleAddTag = () => {
    if (tagInput.trim() && !tags.includes(tagInput.trim())) {
      setTags([...tags, tagInput.trim()]);
      setTagInput('');
    }
  };

  const handleRemoveTag = (tagToRemove) => {
    setTags(tags.filter(tag => tag !== tagToRemove));
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      handleAddTag();
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    // Validate
    if (!title.trim() || !content.trim()) {
      setError('Title and content are required');
      return;
    }

    const postData = {
      title: title.trim(),
      content: content.trim(),
      tags: tags
    };

    // ✅ Use React Query mutation
    createPost(postData, {
      onSuccess: (data) => {
        console.log('Post created:', data);
        // Reset form
        setTitle('');
        setContent('');
        setTags([]);
        setTagInput('');
        // Call onSuccess callback if provided
        if (onSuccess) onSuccess(data);
      },
      onError: (err) => {
        setError(err.message || 'Failed to create post');
      },
    });
  };

  return (
    <form onSubmit={handleSubmit} className="create-post-form">
      <h2>Create New Post</h2>

      {error && <div className="error-message">{error}</div>}

      {/* Title Input */}
      <div className="form-group">
        <label htmlFor="title">Title</label>
        <input
          id="title"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Enter post title..."
          required
          disabled={isPending}
        />
      </div>

      {/* Content Textarea */}
      <div className="form-group">
        <label htmlFor="content">Content</label>
        <textarea
          id="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="Write your post content..."
          rows={5}
          required
          disabled={isPending}
        />
      </div>

      {/* Tags Section */}
      <div className="form-group">
        <label>Tags</label>
        <div className="tag-input-container">
          <input
            type="text"
            value={tagInput}
            onChange={(e) => setTagInput(e.target.value)}
            onKeyDown={handleKeyDown}
            placeholder="Add a tag (press Enter)"
            disabled={isPending}
          />
          <button
            type="button"
            onClick={handleAddTag}
            disabled={isPending}
          >
            Add
          </button>
        </div>
        <div className="tags-display">
          {tags.length === 0 ? (
            <span className="no-tags">No tags added</span>
          ) : (
            tags.map(tag => (
              <span key={tag} className="tag-item">
                #{tag}
                <button
                  type="button"
                  onClick={() => handleRemoveTag(tag)}
                  disabled={isPending}
                >
                  ×
                </button>
              </span>
            ))
          )}
        </div>
      </div>

      {isPending && <p className="loading-text">Creating post...</p>}

      <button
        type="submit"
        disabled={isPending || isLoading}
        className="submit-button"
      >
        {isPending ? 'Publishing...' : 'Publish Post'}
      </button>
    </form>
  );
}
