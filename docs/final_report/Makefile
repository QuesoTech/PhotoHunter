TARGET=report

all:
	pdflatex $(TARGET)
	pdflatex $(TARGET)
	pdflatex $(TARGET)

#dot:
#	dot ../midterm_presentation/figs/er.dot -Tpng -o ../midterm_presentation/figs/er.png

spell:
	aspell -ttex -c $(TARGET).tex

clean:
	rm -f $(TARGET).pdf *.blg *.aux *.bbl *.brf *.log *.nav *.snm *.toc *.out
