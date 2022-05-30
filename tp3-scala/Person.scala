import scala.io.StdIn.{readLine,readInt,readBoolean}


class Library {
  // Nothing here
}

class Document {
  var docNumber: Int = 0
  var publishYear: Int = 0
  var title: String = ""

  def getDocNumber(): Unit = {
    return docNumber
  }
}

class Book extends Library {
  var authorName: String = ""
  var numberOfPages: Int = 100

  def getNumberOfPages(): Unit = {
    return numberOfPages
  }

  def getAuthor(): Unit = {
    return authorName
  }
}

class Article extends Library {
  var journalName: String = ""
  var publishYear: Int = 0
  var volume: Int = 0
  private var month: String = ""

  def getJournal(): Unit = {
    return journalName
  }

  def getPublishYear(): Unit = {
    return publishYear
  }

  def showJournal(): Unit = {
    println("Le journale est yyy")
  }
}

class Dictionary extends Library {
  var longSource: String = ""
  var longCible: String = ""
  var numberOfWords: Int = 0

  def getLongSource(): Unit = {
    return longSource
  }

  def betLongCible(): Unit = {
    return longCible
  }

  def getNumberOfWords(): Unit = {
    return numberOfWords
  } 
}

class Roman extends Book {
  var Price: Int = 0
}

class Manual extends Book {
  var scholarityLevel: String = ""
}

object main extends App {
  // instantiation examples
  var library: Library = new Library
  var book: Book = new Book
  var document: Document = new Document
  var article: Article = new Article
  var dictionary: Dictionary = new Dictionary
  var roman: Roman = new Roman
  var manual: Manual = new Manual

  // method calls example
  book.getNumberOfPage()
  
  // and other method calls if you wish to, but I ll just stop here,
  // feel free to try out the code, thank you Mrs Zouaoui.
}