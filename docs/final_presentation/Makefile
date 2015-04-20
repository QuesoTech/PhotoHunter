TARGET=pres
FIGS=figs

all: spell pdf clean

pdf: dot
	pdflatex $(TARGET)
#	bibtex $(TARGET)
	pdflatex $(TARGET)
#	pdflatex $(TARGET)

dot:
	dot $(FIGS)/er.dot -Tpng -o $(FIGS)/er.png

spell:
	aspell -ttex -c $(TARGET).tex

clean:
	rm -f *.blg *.aux *.bbl *.brf *.log *.nav *.snm *.toc *.out
