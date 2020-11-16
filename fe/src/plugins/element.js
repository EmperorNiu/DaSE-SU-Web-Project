import Vue from 'vue'
import {
  Button, Form, FormItem, Input, Message,
  Menu, MenuItem, MenuItemGroup, Submenu,
  Container, Header, Main, Aside, Col, Row,
  Breadcrumb, BreadcrumbItem, Card, Drawer,
  Radio, RadioGroup, Avatar, Tag, Calendar, Pagination,
  Backtop, Table, TableColumn, Divider, Upload, Switch,
  Notification, Image, Timeline, TimelineItem, Tooltip,
  Dialog, Select, Option, Alert, Link, Carousel, CarouselItem
} from 'element-ui'

Vue.use(Switch)
Vue.use(Pagination)
Vue.use(Button)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Carousel)
Vue.use(CarouselItem)
Vue.use(Dialog)
Vue.use(Select)
Vue.use(Option)
Vue.use(Image)
Vue.use(Timeline)
Vue.use(TimelineItem)
Vue.use(Input)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)
Vue.use(Container)
Vue.use(Header)
Vue.use(Main)
Vue.use(Aside)
Vue.use(Breadcrumb)
Vue.use(BreadcrumbItem)
Vue.use(Card)
Vue.use(Col)
Vue.use(Row)
Vue.use(Drawer)
Vue.use(Divider)
Vue.use(Radio)
Vue.use(RadioGroup)
Vue.use(Avatar)
Vue.use(Tag)
Vue.use(Calendar)
Vue.use(Backtop)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Upload)
Vue.use(Carousel)
Vue.use(CarouselItem)
// Vue.use(Notification)
Vue.use(Alert)
Vue.use(Link)
Vue.use(Tooltip)
Vue.prototype.$message = Message
Vue.prototype.$notify = Notification
