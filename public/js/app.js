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
                        input.name = 'todo';
                        input.value = text;

                        let index = todos[j].querySelector('[name="index"]').getAttribute('value');

                        console.log(todos[j].querySelector('[name="index"]'));
                        console.log(`index: ${index}`);

                        let inputIndex = document.createElement('input');

                        inputIndex.setAttribute('value', index);
                        inputIndex.name = 'index'
                        inputIndex.setAttribute('hidden', true);
    
                        todos[i].innerHTML = '';
                        todos[i].appendChild(input);
                        todos[i].appendChild(inputIndex);
                    }
                } else {
                    if (!label) {
                        let input = todos[j].querySelector('input');

                        let text = input.value;

                        let label = document.createElement('label');

                        label.classList.add('col-form-label');
                        label.innerText = text;

                        let index = todos[j].querySelector('[name="index"]').getAttribute('value');

                        let inputIndex = document.createElement('input');

                        inputIndex.setAttribute('value', index);
                        inputIndex.name = 'index'
                        inputIndex.setAttribute('hidden', true);

                        todos[j].innerHTML = '';
                        todos[j].appendChild(label);
                        todos[j].appendChild(inputIndex);
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

                    let index = todos[i].querySelector('[name="index"]').getAttribute('value');

                    let inputIndex = document.createElement('input');

                    inputIndex.setAttribute('value', index);
                    inputIndex.name = 'index'
                    inputIndex.setAttribute('hidden', true);

                    todos[i].innerHTML = '';
                    todos[i].appendChild(label);
                    todos[i].appendChild(inputIndex);
                }
            }
        }
    })
});