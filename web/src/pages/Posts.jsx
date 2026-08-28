// pages/Posts.jsx
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchPosts } from '../hooks/Post';
import CreatePost from '../components/CreatePost';
import PostCard from '../components/PostCard';

export default function PostsPage() {
  const [showCreate, setShowCreate] = useState(false);

  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['posts'],
    queryFn: () => fetchPosts(),
    staleTime: 5 * 60 * 1000,
  });

  if (isLoading) return <div>Loading posts...</div>;
  if (isError) return <div>Error: {error.message}</div>;

  const posts = data?.allPosts || data?.posts || data || [];
  console.log(posts)
  return (
    <div className="posts-page">
      <div className="header">
        <h1>Posts</h1>
        <button onClick={() => setShowCreate(!showCreate)}>
          {showCreate ? 'Cancel' : '+ New Post'}
        </button>
      </div>

      {showCreate && (
        <CreatePost onSuccess={() => {
          setShowCreate(false);
          refetch(); // Refresh posts
        }} />
      )}

      <div className="posts-grid">
        {posts.map(post => (
          <PostCard key={post.id} post={post} />
        ))}
      </div>
    </div>
  );
}
