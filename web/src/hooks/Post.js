import { backend_url } from "../utils/BackendCall";

// GET posts (with optional filters)
export const fetchPosts = async (filters = {}) => {
  try {
    const response = await fetch(`${backend_url}/api/posts`, {
      method: 'GET', // or GET depending on your API
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching posts:", error);
    throw error;
  }
};

// CREATE post
export const createPost = async (postData) => {
  try {
    const response = await fetch(`${backend_url}/api/post`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(postData),
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error creating post:", error);
    throw error;
  }
};
