MAIN = main
#LATEX = xelatex
#LATEX = pdflatex -shell-escape
LATEX = latex
ZIPNAME = aps-zip
ZIPFOLDER = report-zip.d
zipthesefiles = *.tex *.bib *.pdf images *.sty Makefile
maindependencies = $(wildcard *.tex)
maindependencies += $(wildcard *.bib)
maindependencies += $(wildcard *.sty)

$(MAIN).pdf: $(maindependencies) 
	$(LATEX) $(MAIN) && $(LATEX) $(MAIN) && dvips $(MAIN).dvi && ps2pdf $(MAIN).ps
	#$(LATEX) $(MAIN) && bibtex $(MAIN) && $(LATEX) $(MAIN) && $(LATEX) $(MAIN) && dvips $(MAIN).dvi && ps2pdf $(MAIN).ps

clean:
	$(RM) *.aux *.bbl *.dvi *.log *.out *.toc *.blg *.lof *.lot *.ps *.nav *.snm *.vrb *.pdf

zip: $(MAIN).pdf
	if [ ! -d $(ZIPFOLDER) ]; then mkdir $(ZIPFOLDER); fi
	echo "Creating zip file $(ZIPNAME).zip"
	if [ -f $(ZIPNAME).zip ]; then rm -f $(ZIPNAME).zip; fi
	rm -fr $(ZIPFOLDER)/*
	cp -r $(zipthesefiles) $(ZIPFOLDER)
	zip -r $(ZIPNAME).zip $(ZIPFOLDER)
	echo "Done."

show:
	evince $(MAIN).pdf &

.DEFAULT: $(MAIN).pdf
.PHONY: show zip clean
