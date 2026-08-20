import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
export default function Signup() {
  const [formData, SetFormData] = useState({ first_name: '', last_name: '', username: '', email: '', password: '' })
  const [err, setErr] = useState('')
  const navigate = useNavigate()
  const handleSubmit = async (e) => {
    e.preventDefault();
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };
  try {
    const response = await fetch('http://localhost:8080/v1/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData),
    });

    if (!response.ok) {
      const errData = await response.json();
      throw new Error(errData.error || 'Sign up failed');
    }

    navigate('/welcome');
  } catch (err) {
    setErr(err.message);
  }
};
return (
  <form onSubmit={handleSubmit}>
    <div className="sign-box">
      <div className="sign-human-names">
        <span> First Name </span>
        <input type="text" name="first_name" placeholder="john" required />
        <span> Last Name </span>
        <input type="text" name="last_name" placeholder="doe" required />
      </div>
      <span className="sign-titles"> Username</span>
      <input type="text" name="username" placeholder="johndoe" required />
      <span className="sign-titles"> Password </span>
      <input type="password" name="password" placeholder="******" required />
      <button type="submit">Create Account</button>
    </div>
  </form>

)
}
