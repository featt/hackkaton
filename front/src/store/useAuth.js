import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'

export const useAuthState = create(persist((set) => ({
    user: null,
    setUser: (user) => set({ user }),
}), {
    name: 'user',
    storage: createJSONStorage(() => sessionStorage),
}));
  
