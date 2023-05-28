import {
  HomeIcon,
  UserCircleIcon,
  TableCellsIcon,
  BellIcon,
  ArrowRightOnRectangleIcon,
  UserPlusIcon,
} from "@heroicons/react/24/solid";
import { Home, Profile, Tables, Notifications } from "@/pages/dashboard";
import LoginPage from "@/pages/auth/LoginPage";
import CandidatFrom from '@/pages/CandidatFrom'
import RegistrationPage from "@/pages/auth/RegisterPage";
import Lectures from "./pages/dashboard/lectures";
import VacancyFrom from "./pages/dashboard/vacancyForm";

const icon = {
  className: "w-5 h-5 text-inherit",
};

export const routes = [
  {
    layout: "dashboard",
    pages: [
      {
        icon: <HomeIcon {...icon} />,
        name: "Панель управления",
        path: "/home",
        element: <Home />,
      },
      {
        icon: <UserCircleIcon {...icon} />,
        name: "Профиль",
        path: "/profile",
        element: <Profile />,
      },
      {
        icon: <UserCircleIcon {...icon} />,
        name: "Анкета",
        path: "/form",
        element: <CandidatFrom />,
      },
      {
        icon: <TableCellsIcon {...icon} />,
        name: "Таблицы",
        path: "/tables",
        element: <Tables />,
      },
      {
        icon: <TableCellsIcon {...icon} />,
        name: "Лекции",
        path: "/lectures",
        element: <Lectures />,
      },
      {
        icon: <TableCellsIcon {...icon} />,
        name: "Лекции",
        path: "/vacancy",
        element: <VacancyFrom />,
      },
      {
        icon: <BellIcon {...icon} />,
        name: "Уведомления",
        path: "/notifactions",
        element: <Notifications />,
      },
    ],
  },
  {
    title: "",
    layout: "auth",
    pages: [
      {
        icon: <ArrowRightOnRectangleIcon {...icon} />,
        name: "Вход",
        path: "/login",
        element: <LoginPage />,
      },
      {
        icon: <UserPlusIcon {...icon} />,
        name: "Регистрация",
        path: "/register",
        element: <RegistrationPage />,
      },
    ],
  },
];

export default routes;
