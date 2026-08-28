// hooks/useCreatePost.js
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createPost } from './Post';

export const useCreatePost = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createPost,
    onSuccess: (newPost) => {
      queryClient.invalidateQueries({ queryKey: ['posts'] });

      console.log('Post created successfully:', newPost);
    },
    onError: (error) => {
      console.error('Failed to create post:', error);
    },
  });
};
