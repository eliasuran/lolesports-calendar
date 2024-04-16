from mwrogue.esports_client import EsportsClient
from mwrogue.esports_client import EsportsClient


def api_test():
    site = EsportsClient("lol")

    response = site.cargo_client.query(
        tables="ScoreboardGames=SG, Tournaments=T",
        join_on="SG.OverviewPage=T.OverviewPage",
        fields="T.Name, SG.DateTime_UTC, SG.Team1, SG.Team2",
        where="SG.DateTime_UTC >= '2019-08-01 00:00:00'",  # Results after Aug 1, 2019
        limit=50
    )

    print(response)
