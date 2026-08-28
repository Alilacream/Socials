import "../styles/postcards.css"
export default function PostCard(props) {
  const tags = props.post.tags
  // ✅ Properly parse ISO date
  const parseDate = (dateString) => {
    // Handle null/undefined
    if (!dateString) return null;

    // Create date object
    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) {
      console.warn('Invalid date:', dateString);
      return null;
    }

    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  }

  const displayedTags = tags.map((tag, index) => {
    <span key={index} className="tags">{tag}</span>
  })

  return (
    <div className="postcard">
      <h1>{props.post.title}</h1>
      <p className="post-content">{props.post.content}</p>
      <div className="tags-container">
        {displayedTags}
      </div>
      <div className="dates">
        <span>created at {parseDate(props.post.created)}</span>
        {props.created !== props.updated && <span>{parseDate(props.post.updated)}</span>}
      </div>
    </div>
  )
}

