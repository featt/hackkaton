import {
  BanknotesIcon,
  UserPlusIcon,
  UserIcon,
  ChartBarIcon,
} from "@heroicons/react/24/solid";

export const statisticsCardsData = [
  {
    color: "pink",
    icon: UserIcon,
    title: "Всего пользователей",
    value: "2,300",
    footer: {
      color: "text-green-500",
      value: "+3%",
      label: "за месяц",
    },
  },
  {
    color: "green",
    icon: UserPlusIcon,
    title: "Всего откликов",
    value: "3,462",
    footer: {
      color: "text-red-500",
      value: "-2%",
      label: "за день",
    },
  },
  {
    color: "blue",
    icon: BanknotesIcon,
    title: "Релевантные отклики",
    value: "680",
    footer: {
      color: "text-green-500",
      value: "+55%",
      label: "за последний месяц",
    },
  },
  {
    color: "orange",
    icon: ChartBarIcon,
    title: "Нерелевантных откликов",
    value: "520",
    footer: {
      color: "text-green-500",
      value: "+5%",
      label: "за день",
    },
  },
];

export default statisticsCardsData;
