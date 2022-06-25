# Простейшая реализация коносльного рендерерара из древовидной структуры, подобной HTML

![Screenshot_58](https://user-images.githubusercontent.com/67334706/175764324-d611c2eb-0110-4fc5-a9be-1d41d082dc03.png)

Проект состоит из 3 модулей

- HOM (Hand OF Midas) - модуль занимется высчитыванием границ элементов, размеров. Модуль разбит на две фазы. 1-высчитывание размеров элементов, 2- высчитывание границ элементова основываясь на их размерах
- HOMR (Hand OF Midas Renderer) - Модуль занмиется отображением всего на экран, имея размеры и границы элементов
- HOMF (Hand OF Midas Framework) - Простая оболчка над HOM, Представляет из себя фасад для взаимодействия с бибилиотекой, позволяет отлавливать клики на элементах и запускать перерендер всего содержимого 

Проект очень простой и сырой, разрабатывался с целью развлечения и получения навыков, возможно, когда-то будет доведен до ума, и сможет рендерить HTML/CSS в консоли

Так же, хоть все и работает шустро, тут можно выделить подобие браузерному рендеингу. Все разделено на фазы.
Можно делать минимум работы для изменения цвета, либо половину работы при изменении позиционирования только одного элемента (Вынос из основного потока)
