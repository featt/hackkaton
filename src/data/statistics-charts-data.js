import { chartsConfig } from "@/configs";

const ageStatistic = {
  type: "bar",
  height: 220,
  series: [
    {
      name: "Возраст",
      data: [36, 20, 18, 22],
    },
  ],
  options: {
    ...chartsConfig,
    colors: "#fff",
    plotOptions: {
      bar: {
        columnWidth: "16%",
        borderRadius: 5,
      },
    },
    xaxis: {
      ...chartsConfig.xaxis,
      categories: ["<18", "18-25", "26-30", "31-35"],
    },
  },
};

const dailySalesChart = {
  type: "bar",
  height: 220,
  series: [
    {
      name: "Возраст",
      data: [36, 20, 18, 22],
    },
  ],
  options: {
    ...chartsConfig,
    colors: "#fff",
    plotOptions: {
      bar: {
        columnWidth: "16%",
        borderRadius: 5,
      },
    },
    xaxis: {
      ...chartsConfig.xaxis,
      categories: [
      "нет ОР",
      "есть ПД",
      "есть ВД",
      "есть ОР",],
    },
  },
};
 
const questionnaires = {
  type: "bar",
  height: 220,
  series: [
    {
      name: "Стажировок",
      data: [36, 20, 18, 22, 26, 23, 17],
    },
  ],
  options: {
    ...chartsConfig,
    colors: "#fff",
    plotOptions: {
      bar: {
        columnWidth: "16%",
        borderRadius: 5,
      },
    },
    xaxis: {
      ...chartsConfig.xaxis,
      categories: ["HR", "IT", "Соц", "Мед", "Комф", "Право", "Эконом."],
    },
  },
  };

export const statisticsChartsData = [
  {
    color: "blue",
    title: "Статистика по возрасту",
    chart: ageStatistic,
  },
  {
    color: "pink",
    title: "Статистика по опыту работы",
    chart: dailySalesChart,
  },
  {
    color: "green",
    title: "Статистика стажировок",
    chart: questionnaires,
  },
];

export default statisticsChartsData;
