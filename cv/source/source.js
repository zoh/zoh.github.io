'use strict';
/**
 *  @author zoh
 *  @link http://
 *  @class Persone
 *  @language javascript
 */

function Persone () {

  this.name      = 'Захаров Илья';
  this.birthDay  = '17.08.1988';
  this.phone     = '+7 963 1218671';
  this.email     = 'zoh@bk.ru';
  this.gitHub    = 'https://github.com/zoh';

  this.qualities = 'Целеустремлённый специалист c 5+ летним опытом работы в области IT. \n' +
                   'Ищу возможность применить свои знания и навыки в IT сфере, \n' + 
                   'а также участвовать в разработке сложных, интересных  проектов. \n' +
                   'Заинтересован как в профессиональном, так и в личностном росте. \n\n';

  this.qualities+= '* отличный командный и одиночный разработчик \n' +
                   '* гибкий подход к решению сложных задач \n' +
                   '* непрерывное саморазвитие \n' +
                   '* быстрая обучаемость новым технологиям \n' +
                   '* коммуникабельность, легко вписываюсь в любую среду\n' ;

  // Works
  this.works = [
    {
      period  : '01.2015 - present',
      name    : 'Magnova.ru - Сервис совместных закупок',
      position: 'Front-end developer',
      stack   : 'Angular.js, yeoman.io, Jira',
      note    : 'Ответственен за основную страницу сайта и бек-офис. ' +
                'Для бек-офиса используется AngularJs. ' +
                'Общение с Restful jee сервисом. ' +
                'Работа по Scrum, со спринтами.'
    },
    {
      period  : '08.2014 - 01.2015',
      name    : 'Foturist (http://46.36.219.2/)',
      position: 'Fullstack',
      stack   : 'node.js, sails.js, elasticsearch, PostgreSQL'
    },
    {
      period  : '06.2013 - 01.2015',
      name    : 'ООО КИР, (kirkazan.ru)',
      position: 'Ведущий инженер программист',
      stack   : 'Pure Javascript, Dust.js, Backbone.js, Vanilla.js',
      note    : 'Разработка в команде мета-платформы N2O, для создания корпоративных продуктов. '
    },
    {
      period  : '03.2013 - 06.2013',
      name    : 'TrueTipsters (http://truetipsters.ru)',
      position: 'Fullstack developer',
      stack   : 'CoffeeScript, Node.js, PHP, Redis, MongoDB.',
      note    : 'Разработка распределённого back-end на Node.js'
    }, 
    {
      period  : '06.2012 - 03.2013',
      name    : 'ООО Управление информационными проектами',
      position: 'Web Developer',
      stack   : 'JavaScript, CoffeeScript, KnockoutJS, WebSocket, PHP5',
      note    : 'Для проекта государственных электронных услуг РТ разрабатывал новую версию интерфейса для Инфоматов'
    },
    {
      period  : '01.2012 - 05.2012',
      name    : 'БАРС Груп, ООО (Казань,  www.bars-open.ru)',
      position: 'Web developer',
      stack   : 'JavaScript, ExtJS 4, Python, PostgreSQL',
      note    : 'Создание корпоративной мета-платформы на ExtJS4 + Python работающая по технологии REST для проектов гос. структур.'
    },
    {
      period  : '06.2011 - 01.2012',
      name    : 'ООО "Флашелс"  (Казань, http://flashells.ru/)',
      position: 'старший PHP разработчик',
      stack   : 'PHP5, Yii, MySQL',
      note    : 'Отвечал за backend и всю архитектуру данного проекта.'
    }
  ];


  // Период работы с некоторыми технологиями.
  this.experience = {
    PHP       : '---*******************------********------------------------------',
    MySQL     : '-----******************************-------------------------------',
    JavaScript: '----------------------********************************************',
    NodeJS    : '-------------------------------------******************-----------',
    PostgreSQL: '------------------------------------------------******************',
    Scala     : '-------------------------------------------------------***********',
    Golang    : '-------------------------------------------------------***********',
    Years     : '---2010---/---2011---/---2012---/---2013---/---2014---/---2015---/'
  };


  this.languages = [
    {
      name : 'English',
      level: 'Pre-Intermediate'
    },
    {
      name : 'Русский',
      level: 'родной'
    }
  ];


  // Education
  this.education = {
    yearGraduation: 2010,
    name          : 'КГТУ им Туполева (КАИ)',
    faculty       : 'Технической кибернетики и информатики/' +
                    'Программное обеспечение вычислительной техники, инженер'
  };

  // в добавку.
  this.KeySkills = [
    'CoffeeScript', 'AngularJs', 'JavaScript', 'Backbone.Js', 'Ajax', 'CSS3', 'HTML5',
    'REST', 'JSON', 'TDD', 'BDD', 'MVC', 'MVVM', 'SOA', 'PHP', 'Git', 'MySQL', 'MongoDB', 'Redis',
    'SOLID', 'Elasticsearch', 'Ubuntu', 'CentOs', 'Docker', '...'
  ];

  // achivement, добавить грамоту
  this.achivements = [
    {
      from        : 'ООО УИП',
      achivements : 'благодарственное письмо от директора ООО: "За современный подход в разработке и проектировании ПО"'
    }
  ];
}