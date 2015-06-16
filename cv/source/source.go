package main

import (
	"fmt"
)

func main() {
	p := &Persone{
		Name:     "Захаров Илья",
		BirthDay: "17.08.1988",
		Phone:    "+7 963 1218671",
		Email:    "zoh@bk.ru",
		GitHub:   "https://github.com/zoh",
		Qualities: `Целеустремлённый специалист c 5+ летним опытом работы в области IT.
                    Ищу возможность применить свои знания и навыки в IT сфере,
                    а также участвовать в разработке сложных, интересных  проектов.
                    Заинтересован как в профессиональном, так и в личностном росте.

                    * отличный командный и одиночный разработчик
                    * гибкий подход к решению сложных задач
                    * непрерывное саморазвитие
                    * быстрая обучаемость новым технологиям
                    * коммуникабельность, легко вписываюсь в любую среду`,

		Works: []Work{
			Work{
				Period:   "01.2015 - present",
				Name:     "Magnova.ru - Сервис совместных закупок",
				Position: "Front-end developer",
				Stack:    []string{"Angular.js", "yeoman.io", "Jira"},
				Note: `Ответственен за основную страницу сайта и бек-офис.
                		   Для бек-офиса используется AngularJs.
                		   Общение с Restful jee сервисом.
                		   Работа по Scrum, со спринтами.`,
			},
			Work{
				Period:   "08.2014 - 01.2015",
				Name:     "Foturist (http://46.36.219.2/)",
				Position: "Fullstack",
				Stack:    []string{"node.js", "sails.js", "elasticsearch", "PostgreSQL"},
			},
			Work{
				Period:   "06.2013 - 01.2015",
				Name:     "ООО КИР, (kirkazan.ru)",
				Position: "Ведущий инженер программист",
				Stack:    []string{"Pure Javascript", "Dust.js", "Backbone.js", "Vanilla.js"},
				Note:     "Разработка в команде мета-платформы N2O, для создания корпоративных продуктов.",
			},
			Work{
				Period:   "03.2013 - 06.2013",
				Name:     "TrueTipsters (http://truetipsters.ru)",
				Position: "Fullstack developer",
				Stack:    []string{"CoffeeScript", "Node.js", "PHP", "Redis", "MongoDB"},
				Note:     "Разработка распределённого back-end на Node.js",
			},
			Work{
				Period:   "06.2012 - 03.2013",
				Name:     "ООО Управление информационными проектами",
				Position: "Web Developer",
				Stack:    []string{"JavaScript", "CoffeeScript", "KnockoutJS", "WebSocket", "PHP5"},
				Note: `Для проекта государственных электронных услуг РТ разрабатывал новую 
						версию интерфейса для Инфоматов`,
			},
			Work{
				Period:   "01.2012 - 05.2012",
				Name:     "БАРС Груп, ООО (Казань,  www.bars-open.ru)",
				Position: "Web developer",
				Stack:    []string{"JavaScript", "ExtJS 4", "Python", "PostgreSQL"},
				Note: `Создание корпоративной мета-платформы на ExtJS4 + Python работающая по технологии REST
				 		для проектов гос. структур.`,
			},
			Work{
				Period:   "06.2011 - 01.2012",
				Name:     "ООО \"Флашелс\"  (Казань, http://flashells.ru/)",
				Position: "старший PHP разработчик",
				Stack:    []string{"PHP5", "Yii", "MySQL"},
				Note:     "Отвечал за backend и всю архитектуру данного проекта.",
			},
		},
		Experience: map[string]string{
			"PHP":        "---*******************------********------------------------",
			"MySQL":      "-----******************************-------------------------",
			"JavaScript": "----------------------**************************************",
			"NodeJS":     "-------------------------------------******************-----",
			"PostgreSQL": "------------------------------------------------************",
			"Scala":      "-------------------------------------------------------*****",
			"Golang":     "-------------------------------------------------------*****",
			"Years":      "---2010---/---2011---/---2012---/---2013---/---2014---/---2015---",
		},
		Languages: []Language{
			Language{"Русский", "родной"},
			Language{"English", "Pre-Intermediate"},
		},
		Education: Education{
			2010,
			"КГТУ им Туполева (КАИ)",
			`Технической кибернетики и информатики
			 Программное обеспечение вычислительной техники, инженер`,
		},
		KeySkills: []string{
			"CoffeeScript", "AngularJs", "JavaScript", "Backbone.Js", "Ajax", "CSS3", "HTML5",
			"REST", "JSON", "TDD", "BDD", "MVC", "MVVM", "SOA", "PHP", "Git", "MySQL", "MongoDB", "Redis",
			"SOLID", "Golang", "Akka", "Elasticsearch", "Ubuntu", "CentOS", "Docker", "...",
		},
		Achivements: []Achivement{Achivement{
			From:       "ООО УИП",
			Achivement: "благодарственное письмо от директора ООО: \"За современный подход в разработке и проектировании ПО\"",
		}},
	}
	fmt.Println(p)
}

type Email string

type Work struct {
	Period   string
	Name     string
	Position string
	Stack    []string
	Note     string
}
type Language struct {
	Name  string
	Level string
}
type Education struct {
	YearGraduation int
	Name           string
	Faculty        string
}
type Achivement struct {
	From       string
	Achivement string
}
type Persone struct {
	Name        string
	BirthDay    string
	Phone       string
	Email       Email
	GitHub      string
	Qualities   string
	Works       []Work
	Experience  map[string]string
	Languages   []Language
	Education   Education
	KeySkills   []string
	Achivements []Achivement
}
