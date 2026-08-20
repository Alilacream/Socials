export default function Signin() {
  return (
    <div className="sign-box">
      <span className="sign-titles"> Username</span>
      <input type="text" name="username" placeholder="johndoe" />
      <span className="sign-titles"> Password </span>
      <input type="password" name="password" placeholder="******" />
    </div>
  )
}
