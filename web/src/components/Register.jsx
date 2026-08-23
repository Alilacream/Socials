import { useState } from "react"


export default function LoginPage() {
  const [data, setData] = useState({ first_name: "", last_name: "", username: "", email: "", password: "" })
  const [responed, setResponed] = useState("")
  const handleOnChange = (event) => {
    const { name, value } = event.target
    setData(prev => ({
      ...prev,
      [name]: value
    }));

  }

  const handleSubmit = async (data, event) => {
    event.preventDefault();
    const response = await fetch("localhost:8080/v1/register", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),

    })
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const result = response.json();
    setResponed(result)
  }

  return (
    <form onSubmit={handleSubmit}>

      <span className="sign-title"> First Name</span>
      <input type="text" name="first_name" placeholder="john" onChange={handleOnChange} />

      <span className="sign-title"> Last Name</span>
      <input type="text" name="last_name" placeholder="john" onChange={handleOnChange} />

      <span className="sign-title"> Username </span>
      <input type="text" name="username" placeholder="johndoe" onChange={handleOnChange} />

      <span className="sign-title"> email </span>
      <input type="email" name="email" placeholder="johndoe" onChange={handleOnChange} />

      <span className="sign-title"> Password </span>
      <input type="password" name="password" placeholder="*****" min={8} onChange={handleOnChange} />

      <button type="submit" className="button">Log In</button>
    </form>
  )
}
