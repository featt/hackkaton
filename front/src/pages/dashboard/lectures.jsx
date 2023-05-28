import {
    Card,
    CardHeader,
    CardBody,
    CardFooter,
    Typography,
    Button
  } from "@material-tailwind/react";
   
function LectureCard({url, title, desc}) {
    return (
      <Card className="mt-6 w-96">
        <CardHeader color="blue-gray" className="relative h-56">
         <iframe className="w-full h-full" src={url} />
        </CardHeader>
        <CardBody>
          <Typography variant="h5" color="blue-gray" className="mb-2">
            {title}
          </Typography>
          <Typography>
            {desc}
          </Typography>
        </CardBody>
        <CardFooter className="pt-0">
          <Button>Открыть</Button>
        </CardFooter>
      </Card>
    );
  }

const Lectures = () => {
  const lectures = [
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
    {
        url: 'https://www.youtube.com/embed/AXkOli6BsBY',
        title: 'Асинхронное программирование',
        desc: 'Сегодня расскажим про асихронное программирование в JavaScript'
    },
  ]
  return (
    <div className="mt-12">
      <div className="mb-12 grid gap-y-10 gap-x-6 md:grid-cols-2 xl:grid-cols-4">
        {lectures.map(l => <LectureCard {...l} />)}
      </div>
    </div>
  );
};


export default Lectures