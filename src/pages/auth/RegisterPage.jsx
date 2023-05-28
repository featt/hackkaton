import instance from "@/configs/axiosConfig";
import { useAuthState } from "@/store/useAuth";
import Toast from "@/widgets/Toast";
import React, { useState } from "react";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";


const RegistrationPage = () => {
  const navigate = useNavigate();
  const setUser = useAuthState(state => state.setUser);
  const [open, setOpen] = useState(false)
  const [err, setErr] = useState('')
  
  const reg = useMutation(newReg => {
    return instance.post('/users',newReg)
  })

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = new FormData(e.currentTarget)
    const email = data.get('email')
    const name = data.get('name')
    const last_name = data.get('last_name')
    const password = data.get('password')
    const retryPass = data.get('retryPass')
    if(password !== retryPass) {
      setErr('Пароли не совпадают!')
      setOpen(true)
    }

    await reg.mutateAsync({
      email,
      password,
      name,
      last_name
    })

    navigate('/auth/login');
  };
  
  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-center sm:py-12">
    <div className="p-10 xs:p-0 mx-auto md:w-full md:max-w-md">
      <h1 className="font-bold text-center text-2xl mb-5">Регистрация</h1> 
      <div className="bg-white shadow w-full rounded-lg divide-y divide-gray-200">
        <form onSubmit={handleSubmit} className="px-5 py-7">
          <label className="font-semibold text-sm text-gray-600 pb-1 block">Имя</label>
          <input name="name" type="text" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full" />
          <label className="font-semibold text-sm text-gray-600 pb-1 block">Фамилия</label>
          <input name="last_name" type="text" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full" />
          <label className="font-semibold text-sm text-gray-600 pb-1 block">Почта</label>
          <input name="email" type="email" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full" />
          <label className="font-semibold text-sm text-gray-600 pb-1 block">Пароль</label>
          <input name="password" type="password" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full" />
          <label className="font-semibold text-sm text-gray-600 pb-1 block">Повторите пароль</label>
          <input name="retryPass" type="password" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full" />
          <button type="button" className="transition duration-200 bg-blue-500 hover:bg-blue-600 focus:bg-blue-700 focus:shadow-sm focus:ring-4 focus:ring-blue-500 focus:ring-opacity-50 text-white w-full py-2.5 rounded-lg text-sm shadow-sm hover:shadow-md font-semibold text-center inline-block">
              <button type="submit" className="inline-block mr-2">Регистрация</button>
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" className="w-4 h-4 inline-block">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
          </button>
        </form>           
          <div className="py-5">
          <div className="grid grid-cols-2 gap-1">
            <div className="text-center sm:text-left whitespace-nowrap">
              <button className="transition duration-200 mx-5 px-5 py-4 font-normal text-sm rounded-lg text-gray-500">
                  <span className="inline-block ml-1">Есть аккаунта?</span>
              </button>
            </div>
            <div className="text-center sm:text-right whitespace-nowrap">
              <button className="transition duration-200 mx-5 px-5 py-4 cursor-pointer font-normal text-sm rounded-lg text-gray-500 hover:bg-gray-100 focus:outline-none focus:bg-gray-200 focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50 ring-inset">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" className="w-4 h-4 inline-block align-text-bottom	">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192l-3.536 3.536M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-5 0a4 4 0 11-8 0 4 4 0 018 0z" />
                  </svg>
                  <span onClick={() => navigate('/auth/login')} className="inline-block ml-1">Войти</span>
              </button>
            </div>
          </div>
        
        </div>
      </div>          
      {open && <Toast error={err} open={open} setOpen={setOpen} />}
    </div>
  </div>
  );
};

export default RegistrationPage;
