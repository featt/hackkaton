import PropTypes from "prop-types";
import { NavLink } from "react-router-dom";
import { XMarkIcon } from "@heroicons/react/24/outline";
import { Button, IconButton, Typography } from "@material-tailwind/react";
import { useMaterialTailwindController, setOpenSidenav } from "@/context";
import { useAuthState } from "@/store/useAuth";
import {
  HomeIcon,
  UserCircleIcon,
  TableCellsIcon,
  BellIcon,
  ArrowRightOnRectangleIcon,
  UserPlusIcon,
} from "@heroicons/react/24/solid";

const icon = {
  className: "w-5 h-5 text-inherit",
};

const candidateView = (sidenavType, sidenavColor) => (
  <>
  <li>
    <NavLink to='/dashboard/profile'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <UserCircleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Профиль
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/form'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <UserCircleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Анкета
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/notifactions'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <BellIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Уведомления
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/auth/login'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <ArrowRightOnRectangleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Выйти
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  </>
);

const curatorView = (sidenavType, sidenavColor) => (
  <>
  <li>
    <NavLink to='/dashboard/profile'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <UserCircleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Профиль
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/home'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <HomeIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Панель управления
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/tables'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <TableCellsIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Таблицы
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/notifactions'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <BellIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Уведомления
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/auth/login'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <ArrowRightOnRectangleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Выйти
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  </>
);

const internView = (sidenavType, sidenavColor) => (
  <>
  <li>
    <NavLink to='/dashboard/profile'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <UserCircleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Профиль
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/lectures'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <HomeIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Лекции
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/dashboard/notifactions'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <BellIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Уведомления
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/auth/login'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <ArrowRightOnRectangleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Выйти
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  </>
);

const cadrView = (sidenavType, sidenavColor) => (
  <>
  <li>
    <NavLink to='/dashboard/profile'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <UserCircleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Профиль
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  <li>
    <NavLink to='/dashboard/vacancy'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <HomeIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Вакансия
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/dashboard/notifactions'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <BellIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Уведомления
          </Typography>
        </Button>
      )}
    </NavLink>
  </li> 
  <li>
    <NavLink to='/auth/login'>
      {({ isActive }) => (
        <Button
          variant={isActive ? "gradient" : "text"}
          color={
            isActive
              ? sidenavColor
              : sidenavType === "dark"
              ? "white"
              : "blue-gray"
          }
          className="flex items-center gap-4 px-4 capitalize"
          fullWidth
        >
          <ArrowRightOnRectangleIcon {...icon} />
          <Typography color="inherit" className="font-medium capitalize">
            Выйти
          </Typography>
        </Button>
      )}
    </NavLink>
  </li>
  </>
);


const ViewForRole = ({role, sidenavType, sidenavColor}) => {
  switch(role) {
    case 'candidate':
      return curatorView(sidenavType, sidenavColor)
    case 'curator':
      return curatorView(sidenavType, sidenavColor)
    case 'intern':
      return internView(sidenavType, sidenavColor)
    case 'cadr':
      return cadrView(sidenavType, sidenavColor)
  }
}

export function Sidenav() {
  const user = useAuthState((state) => state.user);
  const [controller, dispatch] = useMaterialTailwindController();
  const { sidenavColor, sidenavType, openSidenav } = controller;
  const sidenavTypes = {
    dark: "bg-gradient-to-br from-blue-gray-800 to-blue-gray-900",
    white: "bg-white shadow-lg",
    transparent: "bg-transparent",
  };

  return (
    <aside
      className={`${sidenavTypes[sidenavType]} ${
        openSidenav ? "translate-x-0" : "-translate-x-80"
      } fixed inset-0 z-50 my-4 ml-4 h-[calc(100vh-32px)] w-72 rounded-xl transition-transform duration-300 xl:translate-x-0`}
    >
      <div
        className={`relative border-b ${
          sidenavType === "dark" ? "border-white/20" : "border-blue-gray-50"
        }`}
      >
        <IconButton
          variant="text"
          color="white"
          size="sm"
          ripple={false}
          className="absolute right-0 top-0 grid rounded-br-none rounded-tl-none xl:hidden"
          onClick={() => setOpenSidenav(dispatch, false)}
        >
          <XMarkIcon strokeWidth={2.5} className="h-5 w-5 text-white" />
        </IconButton>
      </div>
      <div className="m-4">
        
          <ul className="mb-4 flex flex-col gap-1">            
            <ViewForRole role={user.role} sidenavType={sidenavType} sidenavColor={sidenavColor} />
          </ul>
      
      </div>
    </aside>
  );
}

Sidenav.propTypes = {
  brandImg: PropTypes.string,
  brandName: PropTypes.string,
  routes: PropTypes.arrayOf(PropTypes.object).isRequired,
};

Sidenav.displayName = "/src/widgets/layout/sidnave.jsx";

export default Sidenav;
