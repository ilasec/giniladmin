"use client"

import { useState, useEffect } from "react"
import { useRouter } from "next/navigation"
import { fetchCaptcha, login } from "@/services/api"

export default function LoginPage() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [captcha, setCaptcha] = useState("")
  const [captchaId, setCaptchaId] = useState("")
  const [captchaImage, setCaptchaImage] = useState("")
  const [error, setError] = useState("")
  const router = useRouter()

  const loadCaptcha = async () => {
    try {
      const { result } = await fetchCaptcha()
      setCaptchaId(result.CaptchaId)
      setCaptchaImage(result.PicPath)
    } catch (err) {
      setError("获取验证码失败")
    }
  }

  useEffect(() => {
    const token = localStorage.getItem("token")
    if (token) {
      router.push("/")
    }
    loadCaptcha()
  }, [])

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      const { result } = await login({
        username,
        password,
        captcha,
        captchaId
      });
      // 处理登录成功
      localStorage.setItem("token", result.token);
      localStorage.setItem("isAuthenticated", "true");
      router.push("/");
    } catch (err) {
      setError("登录失败");
      loadCaptcha(); // 刷新验证码
    }
  }

  return (
    <div className="flex h-screen">
      {/* 左侧背景图片 */}
      <div className="hidden lg:block lg:w-1/2">
        <img
          src="/images/login.png"
          alt="Login background"
          className="h-full w-full object-cover"
        />
      </div>

      {/* 右侧登录表单 */}
      <div className="flex w-full items-center justify-center lg:w-1/2">
        <form onSubmit={handleLogin} className="w-96 space-y-4 rounded-lg bg-white p-8 shadow-lg">
          <h1 className="text-2xl font-bold text-center mb-6">登录</h1>
          {error && <div className="text-red-500 text-sm">{error}</div>}
        <div>
          <label className="block text-sm font-medium text-gray-700">用户名</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">密码</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">验证码</label>
          <div className="flex gap-2">
            <input
              type="text"
              value={captcha}
              onChange={(e) => setCaptcha(e.target.value)}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
            />
            {captchaImage && (
              <img
                src={captchaImage}
                alt="验证码"
                className="mt-1 h-10 cursor-pointer"
                onClick={loadCaptcha}
              />
            )}
          </div>
        </div>
        <button
          type="submit"
          className="w-full rounded-md bg-blue-600 px-4 py-2 text-white hover:bg-blue-700"
        >
          登录
        </button>
      </form>
      </div>
    </div>
  )
}