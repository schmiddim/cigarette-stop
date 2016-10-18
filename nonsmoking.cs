using System;
/**
 * example: mono stopsmoking.exe 2016-09-19 30 0,285714286
 *
 *
 */
namespace stopsmoking
{
	class MainClass
	{
		static String GetSmokingString (DateTime dateStopped, int cigartettesPerDay, double pricePerCigarette)
		{
			DateTime now = DateTime.Now;
			TimeSpan daysWithoutSmoking = now.Subtract (dateStopped);
			int cigarettesNotSmoked = (int)(cigartettesPerDay * daysWithoutSmoking.TotalDays);
			double moneySaved = Math.Round (pricePerCigarette * cigarettesNotSmoked, 2);
			return 	String.Format ("You stopped smoking {0} days ago. Saved {1} Euro and you didnt't smoked {2} cigarettes", (int)daysWithoutSmoking.TotalDays, moneySaved, cigarettesNotSmoked);
		}

		public static void Main (string[] args)
		{
			if (args.Length != 3) {
				System.Console.WriteLine ("Usage smoke-stop (Y-m-D) [cigarettes per day] [price / cigarette] - use the locale delimiter for doubles!");
				return;
			}

			string[] dateParts = args [0].Split ('-');
			Console.WriteLine (GetSmokingString (new DateTime (int.Parse (dateParts [0]), int.Parse (dateParts [1]), int.Parse (dateParts [2]))

			                                    , int.Parse (args [1])
			                                    , double.Parse (args [2])
			)
			);
		}
	}
}
