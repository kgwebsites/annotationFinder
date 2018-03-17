# annotationFinder
Executable file which finds all your // TODO , // FIXME , // REFACTOR annotation comments within your directory and appends them to your README.md file.

## ANNOTATIONS
### TODOS:
* [                        if strings.Contains(string(scanner.Text()), "TODO:") </span><span class="cov8" title="1">{](coverage.html)
* [                                        annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// TODO: ")})](coverage.html)
* [This one has multiple todos!](sample/multipleAnnotations.js)
* [Increase code coverage. Currently at 63.3%](sample/repoAnnotations)
* [Add inclusion and exclusion flags to specify directories to include / exclude](sample/repoAnnotations)
* [This is a todo item](sample/sample.js)
* [foo](testdata/file1)
* [//TODO: foo](testdata/foo)

### FIXME:
* [                        <span class="cov8" title="1">if strings.Contains(string(scanner.Text()), "FIXME:") </span><span class="cov8" title="1">{](coverage.html)
* [                                        annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// FIXME: ")})](coverage.html)
* [This one has multiple todos!](sample/multipleAnnotations.js)
* [bar](testdata/file1)

### REFACTOR:
* [                "REFACTOR": {},](coverage.html)
* [                        <span class="cov8" title="1">if strings.Contains(string(scanner.Text()), "REFACTOR") </span><span class="cov8" title="1">{](coverage.html)
* [                                annotations["REFACTOR"] = append(annotations["REFACTOR"],](coverage.html)
* [                                        annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// REFACTOR: ")})](coverage.html)
* [baz](testdata/file1)

