import {
  Card,
  CardBody,
  CardHeader,
  CardFooter,
  Avatar,
  Typography,
  Tabs,
  TabsHeader,
  Tab,
  Switch,
  Tooltip,
  Button,
} from "@material-tailwind/react";
import {
  HomeIcon,
  ChatBubbleLeftEllipsisIcon,
  Cog6ToothIcon,
  PencilIcon,
} from "@heroicons/react/24/solid";
import { Link } from "react-router-dom";
import { ProfileInfoCard, MessageCard } from "@/widgets/cards";
import { platformSettingsData, conversationsData, projectsData } from "@/data";
import { useAuthState } from "@/store/useAuth";

export function Profile() {
  const user = useAuthState(state => state.user)
  console.log(user);
  return (
    <>
      <div className="relative mt-8 h-10 w-full overflow-hidden rounded-xl  bg-cover	bg-center">
      </div>
      <Card className="mx-3 -mt-16 mb-6 lg:mx-4">
        <CardBody className="p-4">
          <div className="mb-10 flex items-center justify-between gap-6">
            <div className="flex items-center gap-6">             
              <div>
                <Typography variant="h5" color="blue-gray" className="mb-1">
                  {user.name} {user.last_name}
                </Typography>                
              </div>
            </div>
            <div className="w-96">
              <Tabs value="app">
                <TabsHeader>
                <Tab value="settings">
                    <Cog6ToothIcon className="-mt-1 mr-2 inline-block h-5 w-5" />
                    Баллы: 230
                  </Tab>
                  <Tab value="app">
                    <HomeIcon className="-mt-1 mr-2 inline-block h-5 w-5" />
                    Обменять
                  </Tab>
                </TabsHeader>
              </Tabs>
            </div>
          </div>
          <div className="gird-cols-1 mb-12 grid gap-12 px-4 lg:grid-cols-1 xl:grid-cols-2">
            <ProfileInfoCard
              title="Информация о профиле"
              description="Привет, я Степан Плеханов. Решения: Если не можете принять решение, ответ - нет. Если перед вами стоят два одинаково сложных пути, выберите тот, который более болезненен в краткосрочной перспективе (избегание боли создает иллюзию равенства)."
              details={{
                "Имя": "Плеханов Степан",
                "Телефон": "+7 (908) 859 72 32",
                "email": "crayon-d@mail.ru",
                "Локация": "Москва, Россия",
                "Соцсети": (
                  <div className="flex items-center gap-4">
                    <i className="fa-brands fa-facebook text-blue-700" />
                    <i className="fa-brands fa-twitter text-blue-400" />
                    <i className="fa-brands fa-instagram text-purple-500" />
                  </div>
                ),
              }}
              action={
                <Tooltip content="Edit Profile">
                  <PencilIcon className="h-4 w-4 cursor-pointer text-blue-gray-500" />
                </Tooltip>
              }
            />
            <div>
              <Typography variant="h6" color="blue-gray" className="mb-3">
                RoadMap до стажировки
              </Typography>
              <ul className="flex flex-col gap-6">
                {conversationsData.map((props) => (
                  <MessageCard
                    key={props.name}
                    {...props}
                    action={
                      <Button variant="text" size="sm">
                        Начать
                      </Button>
                    }
                  />
                ))}
              </ul>
            </div>
          </div>
          <div className="px-4 pb-4">
            <Typography variant="h6" color="blue-gray" className="mb-2">
              Стажировки
            </Typography>
            <Typography
              variant="small"
              className="font-normal text-blue-gray-500"
            >
              Твои приглашения на стажировку
            </Typography>
            <div className="mt-6 grid grid-cols-1 gap-12 md:grid-cols-2 xl:grid-cols-4">
              {projectsData.map(
                ({ img, title, description, tag, route, members }) => (
                  <Card key={title} color="transparent" shadow={false}>
                    <CardHeader
                      className="mx-0 mt-0 mb-4"
                    >

                    </CardHeader>
                    
                    <CardBody className="py-0 px-1">
                      <Typography
                        variant="small"
                        className="font-normal text-blue-gray-500"
                      >
                        {tag}
                      </Typography>
                      <Typography
                        variant="h5"
                        color="blue-gray"
                        className="mt-1 mb-2"
                      >
                        {title}
                      </Typography>
                      <Link to={route}>
                        <Button variant="outlined" size="sm">
                          Принять
                        </Button>
                      </Link>
                      <Typography
                        variant="small"
                        className="font-normal text-blue-gray-500"
                      >
                        {description}
                      </Typography>
                    </CardBody>
                    <CardFooter className="mt-6 flex items-center justify-between py-0 px-1">
                      <div>
                        {members.map(({ img, name }, key) => (
                          <Tooltip key={name} content={name}>
                            <Avatar
                              src={img}
                              alt={name}
                              size="xs"
                              variant="circular"
                              className={`cursor-pointer border-2 border-white ${
                                key === 0 ? "" : "-ml-2.5"
                              }`}
                            />
                          </Tooltip>
                        ))}
                      </div>
                    </CardFooter>
                  </Card>
                )
              )}
            </div>
          </div>
        </CardBody>
      </Card>
    </>
  );
}

export default Profile;
