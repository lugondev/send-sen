package sen

// ---------- Inline templates (demo) ----------
// In a real project, you can load templates from files or embed them.

// tplWelcome – Welcome email.
const tplWelcome = `
<h2>Hello {{.Name}},</h2>
<p>Welcome to <b>MyService</b>! Explore our amazing features right now.</p>
<p>Best regards,<br/>Team MyService</p>
`

// tplReset – Password reset email.
const tplReset = `
<p>You have requested to reset your password.</p>
<p>Click the link below to continue:</p>
<p><a href="{{.Link}}">{{.Link}}</a></p>
<p>If you didn't request this, please ignore this email.</p>
`

// tplVerificationCode – Verification code email.
const tplVerificationCode = `
<h2>Verification Code</h2>
<p>Your verification code is: <strong>{{.Code}}</strong></p>
<p>This code will expire in 10 minutes.</p>
<p>If you didn't request this code, please ignore this email.</p>
`

// tplWarningLogin – Warning about new login.
const tplWarningLogin = `
<h2>Security Alert: New Login Detected</h2>
<p>We detected a new login to your account from a new location.</p>
<p><strong>Location:</strong> {{.Location}}</p>
<p><strong>Time:</strong> {{.Time}}</p>
<p>If this was you, you can ignore this message. If you didn't log in recently, please secure your account immediately by changing your password.</p>
`
