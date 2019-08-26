document.addEventListener('DOMContentLoaded', function () {
    let todos = document.getElementsByClassName('todo');

    for (let i = 0; i < todos.length; i++) {
        todos[i].addEventListener('click', function () {
            for (let j = 0; j < todos.length; j++) {
                let label = todos[j].querySelector('label');

                if (i === j) {
                    if (label) {
                        let text = label.innerText;
    
                        let input = document.createElement('input');
    
                        input.classList.add('form-control');
                        input.value = text;

                        let index = todos[j].querySelector('span').innerText;

                        let span = document.createElement('span');

                        span.innerText = index;
                        span.setAttribute('hidden', true);
    
                        todos[i].innerHTML = '';
                        todos[i].appendChild(input);
                        todos[i].appendChild(span);
                    }
                } else {
                    if (!label) {
                        let input = todos[j].querySelector('input');

                        let text = input.value;

                        let label = document.createElement('label');

                        label.classList.add('col-form-label');
                        label.innerText = text;

                        let index = todos[j].querySelector('span').innerText;

                        let span = document.createElement('span');

                        span.innerText = index;
                        span.setAttribute('hidden', true);

                        todos[j].innerHTML = '';
                        todos[j].appendChild(label);
                        todos[j].appendChild(span);
                    }
                }
            }
        });
    }

    document.addEventListener('keydown', function (event) {
        if (event.which === 27) {
            for (let i = 0; i < todos.length; i++) {
                let label = todos[i].querySelector('label');

                if (!label) {
                    let input = todos[i].querySelector('input');

                    let text = input.value;

                    let label = document.createElement('label');

                    label.classList.add('col-form-label');
                    label.innerText = text;

                    let index = todos[i].querySelector('span').innerText;

                    let span = document.createElement('span');

                    span.innerText = index;
                    span.setAttribute('hidden', true);

                    todos[i].innerHTML = '';
                    todos[i].appendChild(label);
                    todos[i].appendChild(span);
                }
            }
        }
    })
});